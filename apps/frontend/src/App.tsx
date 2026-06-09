import { useState } from "react";
import Chat from "./components/Chat";

function App() {
	const [count, setCount] = useState(0);

	return (
		<>
			<Chat />
		</>
	);
}

export default App;
