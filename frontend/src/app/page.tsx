import GameCard from "@/components/GameCard";

// const fetchSessions = async (sessionId: string): Promise<unknown> => {
// 	const response = await fetch(`http://localhost/session`, {
// 		method: "GET",
// 		headers: { id: sessionId },
// 	});
// 	const result = await response.json();
// 	return result;
// };
const Home = async () => {
	// const sessions = await fetchSessions("id").then(res => res);

	return (
		<main className="w-full h-screen bg-slate-100 flex justify-center items-center ">
			<GameCard
				id="asdfasdf"
				name="Name1"
				pic="pic"
				isAdmin={true}
				playerCount={3}
			/>
		</main>
	);
};

export default Home;
