import { useEffect, experimental_useEffectEvent } from "react";

function showNotification(message, theme) {
    console.log(`${theme}: ${message}`);
}

function ChatRoom1({ roomId, theme }) {
    // bad
    useEffect(() => {
        const connection = createConnection(serverUrl, roomId);
        connection.on('connected', () => {
            showNotification('Connected!', theme);
        });
        connection.connect();

        return () => {
            connection.disconnect();
        };
    }, [roomId, theme]);
}


// 怎么把theme从connect的effect里面拿出来呢
// 并不想每次有theme改变就连接一次啊
function ChatRoom2({ roomId, theme }) {
    useEffect(() => {
        const connection = createConnection(serverUrl, roomId);
        connection.on('connected', () => {
            showNotification('Connected!', theme);
        });
        connection.connect();

        return () => {
            connection.disconnect();
        };
    }, [roomId, theme]);
}


function ChatRoom3({ roomId, theme }) {

    const [messages, setMessages] = useState([]);
    useEffect(() => {
      const connection = createConnection();
      connection.connect();
      connection.on('message', (receivedMessage) => {
        setMessages([...messages, receivedMessage]);
      });
    }, [messages]);

}

function Chat({ roomId, notificationCount }) {
    const onVisit = useEffectEvent(visitedRoomId => {
      logVisit(visitedRoomId, notificationCount);
    });
  
    useEffect(() => {
      onVisit(roomId);
    }, [roomId]); // ✅ All dependencies declared
    // ...
}


