import { Edit, Play, Share2, Trash } from "lucide-react";
import { useNavigate } from "react-router-dom";

export default function MaticItem({ name, file }) {
  const navigate = useNavigate()
	return (
		<div className="flex p-4 border rounded-md border-slate-700">
			<div className="text-lg font-bold">{name}</div>
			<div className="flex gap-3 ms-auto">
				<button
					title="Edit"
					onClick={() => {
						navigate("/matic/" + file);
					}}
				>
					<Edit></Edit>
				</button>
				<button title="Run">
					<Play></Play>
				</button>
				<button title="Share">
					<Share2></Share2>
				</button>
				<button title="Delete">
					<Trash></Trash>
				</button>
			</div>
		</div>
	);
}
