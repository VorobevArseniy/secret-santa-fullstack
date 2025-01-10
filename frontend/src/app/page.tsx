import SessionCard from "@/components/SessionCard";

export default function Home() {
	return (
		<main className="w-full h-screen flex flex-col bg-orange-50 justify-center items-center">
			<SessionCard
				id="id"
				name="name"
				pic="pic"
				playerCount={12}
				isAdmin={true}
			/>
		</main>
	);
}
