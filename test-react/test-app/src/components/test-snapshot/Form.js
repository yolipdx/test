import { useState } from "react";

export default function Form() {
    const [to, setTo] = useState("Alice");
    const [message, setMessage] = useState("hello");

    function handleSubmit(e) {
        e.preventDefault();

        setTimeout(() => {
            alert(`You said ${message} to ${to}`);
        });
    }

    return (
        <form onSubmit={handleSubmit}>
            <label>
                To: {' '}
                <select
                    value={to}
                    onChange={ e => setTo(e.target.value)}
                >
                    <option value="Alice">Alice</option>
                    <option value="Bob">Bob</option>    
                </select>
            </label>
            <textarea
                placeholder="Message"
                value={message}
                onChange={e => setMessage(e.target.value)}
            />
            <button type="submit">Send</button>
        </form>
    );
}