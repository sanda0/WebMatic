import { useEffect } from "react";
import { useParams } from "react-router-dom";
import Button from "../components/Button";
import { Play, Save, Share2 } from "lucide-react";

export default function Editor() {
	const { name } = useParams();

	useEffect(() => {
		console.log(name);
	});

	return (
		<div className="h-screen p-2">
			<div className="flex p-4 border rounded-md border-slate-700">
				<div className="text-lg font-bold">{name}</div>
				<div className="flex gap-3 ms-auto">
					<button title="Run">
						<Play></Play>
					</button>
					<button title="Save">
						<Save></Save>
					</button>
					<button title="Share">
						<Share2></Share2>
					</button>
				</div>
			</div>
			<div className="mt-4">{name}</div>
		</div>
	);
}
