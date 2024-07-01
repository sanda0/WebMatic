import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Button from "../components/Button";
import { Play, Save, Share2 } from "lucide-react";
import BlockyEmbed from "../components/BlockyEmbed";
import { RunMatic, GetMaticById, SaveXML } from "../../wailsjs/go/main/App";

export default function Editor() {
	const { id } = useParams();
	const [jsonStr, setJsonStr] = useState("");
	const [matic, setMatic] = useState();
	const [xmlStr, setXmlStr] = useState("");

	useEffect(() => {
		console.log(id);
		GetMaticById(parseInt(id)).then((res) => {
			console.log(res);
			if (res.status == 200) {
				setMatic(res.data);
				console.log(res.data);
			}
		});
	}, []);

	function run() {
		RunMatic(jsonStr);
	}

	function savexml() {
		SaveXML(parseInt(id), xmlStr);
	}

	return (
		<div className="h-screen p-2">
			<div className="flex p-4 border rounded-md border-slate-700">
				<div className="text-lg font-bold">
					{matic != null ? matic.name : "-"}
				</div>
				<div className="flex gap-3 ms-auto">
					<button title="Run" onClick={run}>
						<Play></Play>
					</button>
					<button title="Save" onClick={savexml}>
						<Save></Save>
					</button>
					<button title="Share">
						<Share2></Share2>
					</button>
				</div>
			</div>
			{matic && matic.xml_data && (
				<div className="mt-4 h-[80%]">
					<BlockyEmbed
						xml_data={matic.xml_data}
						onChange={(x) => {
							setXmlStr(x);
						}}
						onJsonChange={(j) => {
							setJsonStr(j);
							console.log(j);
						}}
					/>
				</div>
			)}
		</div>
	);
}
