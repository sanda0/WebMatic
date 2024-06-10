import { useEffect } from "react";
import { useParams } from "react-router-dom";
import Button from "../components/Button";
import { Play } from "lucide-react";

export default function Editor() {
	const { name } = useParams();

	useEffect(() => {
		console.log(name);
	});

	return (
		<div className="h-screen p-0">
			<div className="bg-slate-700 h-14">
				<div className="w-10">
					<Button>
						{" "}
						<Play></Play>{" "}
					</Button>
				</div>
			</div>
			{name}
		</div>
	);
}
