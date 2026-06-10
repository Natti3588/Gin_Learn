import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input"

export default function Chat() {
	const [message, setMessage] = useState("");
	const [reply, setReply] = useState("");

	const handleSend = async () => {
		if (!message.trim()) return;
		try {
			const res = await fetch("/api/chat", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ message }),
			});
			const data = await res.json();
			setReply(data.reply);
		} catch (err) {
			// console.error(`送信エラー: ${err}`);
			setReply("エラーが発生しました");
		}
	};

	return (
		<div className="flex flex-col justify-center items-center h-52 gap-4">
			<h1 className="text-lg">Glog</h1>
			<Input
				type="text"
				className="p-2 w-72"
				onChange={(e) => setMessage(e.target.value)}
				placeholder="メッセージを入力"
			/>

			<Button variant="secondary" onClick={handleSend}>
				送信
			</Button>
			<div className="mt-4">
				<strong>応答:</strong> {reply}
			</div>
		</div>
	);
}
