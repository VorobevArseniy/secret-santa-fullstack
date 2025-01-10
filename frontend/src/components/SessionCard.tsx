import Link from "next/link";
import React from "react";

interface GameCardInterface {
	id: string;
	name: string;
	pic: string;
	isAdmin: boolean;
	playerCount: number;
}

const GameCard = ({
	id,
	name,
	pic,
	isAdmin,
	playerCount,
}: GameCardInterface) => {
	const role: string = isAdmin ? "Вы админ" : "Вы участник";
	return (
		<Link
			href={`/${id}`}
			className="w-64 h-52 bg-slate-400  rounded-md flex justify-center items-center shadow-2xl transition transform hover:scale-105 duration-150 cursor-pointer">
			<div className="w-52 h-44 flex flex-col justify-between">
				<div className="flex-1 bg-slate-600">{pic}</div>
				<div className="flex flex-col text-center text-white">
					<p>{name}</p>
					<div className="flex justify-between items-center text-center">
						<p>{playerCount} участников</p>
						<div className="size-1 bg-white rounded-full"></div>
						<p>{role}</p>
					</div>
				</div>
			</div>
		</Link>
	);
};

export default GameCard;
