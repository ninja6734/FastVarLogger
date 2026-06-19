import "./ConnectTab.css"
import { useState } from "react";
import axios from "axios";

function ConnectContent() {

    const [port, setPort] = useState("")
    const [status, setStatus] = useState("")
    const [testData, setTestData] = useState("")

    const handleConnect = async () => {
        try {
            const response = await axios.post('/api/connect', { port, testData });
            setStatus(response.data.message);
        } catch (error) {
            setStatus("Connection failed");
        }
    }
  return (
    <div>
        <h1>Connnect to Robot</h1>
        <p>To connect to the robot, please enter the port number of the robot in the field below. Then click the "Connect" button to establish a connection.</p>
        <input type="text" placeholder="8080" className="input-field" value={port} onChange={(e) => setPort(e.target.value)}/>
        <input type="text" placeholder="Test Data (optional)" className="input-field" value={testData} onChange={(e) => setTestData(e.target.value)}/>
        <button className="border-2 border-black rounded-2xl p-2 mt-4 bg-amber-50 hover:bg-amber-100" onClick={handleConnect}>Connect</button>
        <p>{status}</p>
    </div>
  )
}

export default ConnectContent