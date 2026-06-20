import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

export default function Chat() {
	const [title, setTitle] = useState("");
	const [body, setBody] = useState("");
	const [message, setMessage] = useState("");

	const handleSend = async () => {
		if (!title.trim()) {
			setMessage("タイトルを入力してください");
			return;
		}
		try {
			const res = await fetch("/api/posts/", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ title, body }),
			});

			if (!res.ok) {
				const err = await res.json();
				setMessage(`登録失敗: ${err.error ?? res.status}`);
				return;
			}
			setMessage("登録しました");
		} catch (err) {
			setMessage(err instanceof Error ? err.message : "エラーが発生しました");
		} finally {
			setTitle("");
			setBody("");
		}
	};

	return (
		<div className="flex flex-col justify-center items-center gap-4 p-4">
			<h1 className="text-lg">Glog</h1>
			<Input
				type="text"
				className="p-2 w-72"
				value={title}
				onChange={(e) => setTitle(e.target.value)}
				placeholder="タイトル"
			/>
			<Input
				type="text"
				className="p-2 w-72"
				value={body}
				onChange={(e) => setBody(e.target.value)}
				placeholder="本文"
			/>

			<Button variant="secondary" onClick={handleSend}>
				送信
			</Button>
			{message && <div className="mt-2 text-sm">{message}</div>}
		</div>
	);
}
