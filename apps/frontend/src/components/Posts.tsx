import { useEffect, useState } from "react";

// backend の dto.PostResponse に対応する型
type Post = {
	id: number;
	title: string;
	body: string;
};

export default function Posts() {
	const [posts, setPosts] = useState<Post[]>([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState<string | null>(null);

	useEffect(() => {
		// useEffect のコールバック自体は async にできないため、
		// 中で async 関数を定義して呼び出す（これが定石）
		const fetchPosts = async () => {
			try {
				// vite proxy 経由: /api/posts/ -> http://localhost:8080/posts/
				const res = await fetch("/api/posts/");
				if (!res.ok) {
					throw new Error(`HTTP ${res.status}`);
				}
				const data: Post[] = await res.json();
				setPosts(data);
			} catch (err) {
				setError(err instanceof Error ? err.message : "取得に失敗しました");
			} finally {
				setLoading(false);
			}
		};

		fetchPosts();
	}, []);

	if (loading) {
		return <p className="text-center text-muted-foreground">読み込み中...</p>;
	}

	if (error) {
		return <p className="text-center text-red-500">エラー: {error}</p>;
	}

	return (
		<div className="mx-auto flex max-w-xl flex-col gap-4 p-4">
			<h2 className="text-xl font-bold">投稿一覧</h2>
			{posts.length === 0 ? (
				<p className="text-muted-foreground">投稿がありません</p>
			) : (
				<ul className="flex flex-col gap-3">
					{posts.map((post) => (
						<li key={post.id} className="rounded-lg border p-4">
							<h3 className="font-semibold">{post.title}</h3>
							<p className="text-sm text-muted-foreground">{post.body}</p>
						</li>
					))}
				</ul>
			)}
		</div>
	);
}
