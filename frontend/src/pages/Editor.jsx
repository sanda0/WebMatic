import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Button from "../components/Button";
import { Play, Save, Share2 } from "lucide-react";
import BlockyEmbed from "../components/BlockyEmbed";
import { RunMatic } from "../../wailsjs/go/main/App";

export default function Editor() {
	const { name } = useParams();
	const [jsonStr,setJsonStr] = useState("");

	useEffect(() => {
		console.log(name);
	});

	function run(){
		RunMatic(jsonStr)
	}


	return (
		<div className="h-screen p-2">
			<div className="flex p-4 border rounded-md border-slate-700">
				<div className="text-lg font-bold">{name}</div>
				<div className="flex gap-3 ms-auto">
					<button title="Run" onClick={run}>
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
			<div className="mt-4 h-[80%]">
				<BlockyEmbed onChange={(j)=>{
					console.log("change");
					console.log(j)
					setJsonStr(j)
				}}></BlockyEmbed>
			</div>
		</div>
	);
}
