import keypair from "keypair";
import _ from 'lodash'
import {useRef, useState} from "react";
import EncryptRsa from 'encrypt-rsa';

function App() {
    const [publicKey, setPublicKey] = useState(null);
    const [filesList, setFilesList] = useState([]);
    const [fileContent, setFileContent] = useState(null);
    const [decryptedFileContent, setDecryptedFileContent] = useState(null);

    const firstRef = useRef();
    const secondRef = useRef();

    const onKeyGeneration = () => {
        if (publicKey == null) {
            let pair = keypair();
            let id = _.uniqueId();
            localStorage.setItem("id", id);
            localStorage.setItem('publicKey', pair.public);
            localStorage.setItem('privateKey', pair.private);
            setPublicKey(pair.public);
            fetch('http://localhost:8080/key', {
                method: 'POST',
                mode: 'cors',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    id,
                    payload: pair.public
                })
            }).then(data => console.log(data))
        } else {
            setPublicKey(null)
        }
    }

    const onKeyShow = () => {
        fetch(`http://localhost:8080/key?id=${localStorage.getItem("id")}`, {
            method: 'GET',
            mode: 'cors',
        }).then(response => response.json())
            .then(data => setPublicKey(data.payload))
            .catch((error) => {
                console.error('Error:', error);
                setPublicKey(null);
            });
    }

    const onKeyEnter = (publicKey, privateKey) => {
        let id = _.uniqueId();
        localStorage.setItem("id", id);
        localStorage.setItem('publicKey', publicKey);
        localStorage.setItem('privateKey', privateKey);
        setPublicKey(publicKey);
        fetch('http://localhost:8080/key', {
            method: 'POST',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                id,
                payload: publicKey
            })
        }).then(data => console.log(data))
    }

    const onFileUpload = (event) => {
        const fileReader = new FileReader();
        fileReader.readAsText(event.target.files[0]);
        fileReader.onload = sendFile
    }

    const sendFile = (event) => {
        const encryptRsa = new EncryptRsa();
        const id = localStorage.getItem("id") ?? _.uniqueId();
        fetch('http://localhost:8080/file', {
            method: 'POST',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                id,
                content: encryptRsa.encryptStringWithRsaPublicKey({
                    text: event.target.result,
                    publicKey: localStorage.getItem("publicKey")
                })
            })
        }).then(data => console.log(data))
    }

    const getFiles = () => {
        fetch('http://localhost:8080/file', {
            method: 'GET',
            mode: 'cors',
        }).then(response => response.json())
            .then(data => setFilesList(data));
    };

    useState(() => {
        getFiles();
    });

    const displayContent = (event) => {
        const file = filesList.find(file => file.id === event.target.textContent);
        setFileContent(file.content);
        displayDecryptedContent(file.content);
    }

    const displayDecryptedContent = (content) => {
        const encryptRsa = new EncryptRsa();
        setDecryptedFileContent(
            encryptRsa.decryptStringWithRsaPrivateKey({
                text: content,
                privateKey: localStorage.getItem("privateKey")
            })
        );
    };

    return (
        <>
            <div style={{display: 'flex'}}>
                <div style={{width: '50%', display: 'flex', justifyContent: 'space-evenly'}}>
                    <div>
                        <button onClick={onKeyGeneration}>Generate key</button>
                        {publicKey != null ? <div style={{width: '250px'}}>{publicKey}</div> : null}
                        {localStorage.getItem("id") != null ?
                            <div style={{width: '250px'}}>{`Your id is ${localStorage.getItem("id")}`}</div> : null}
                    </div>
                    <div>
                        <button onClick={onKeyShow}>Show key pair</button>
                    </div>
                    <div style={{display: 'flex', flexDirection: 'column'}}>
                        <input ref={firstRef}/>
                        <input ref={secondRef}/>
                        <button onClick={() => onKeyEnter(firstRef.current.value, secondRef.current.value)}>Enter key
                        </button>
                    </div>
                </div>
                <div style={{width: '50%', display: 'flex'}}>
                    <div>
                        <input onChange={onFileUpload} type="file"/>
                    </div>
                    <div>
                        <table>
                            <thead>
                            <tr>
                                <th scope="col">Id (Filename)</th>
                            </tr>
                            </thead>
                            <tbody style={{cursor: 'pointer'}}>
                            {filesList.map(file => <tr>
                                <th scope="row"><a onClick={displayContent} key={file.id}>{file.id}</a></th>
                            </tr>)}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
            <div style={{display: "flex", justifyContent: 'space-evenly'}}>
                <div style={{width: '50%'}}>
                    {fileContent}
                </div>
                <div style={{width: '50%'}}>
                    {decryptedFileContent}
                </div>
            </div>
        </>
    );
}

export default App;
