import { useState } from "react";
import Input from "../components/Input";
import Button2 from "../components/Button2";
import { Save } from "lucide-react";
import { SaveMatic } from "../../wailsjs/go/main/App";
import { useNavigate } from "react-router-dom";

export default function NewMatic() {
	const [name, setname] = useState("");
	const [author, setAuthor] = useState("");
	const navigate = useNavigate();

	function save(e) {
		console.log(name);
		console.log(author);
		SaveMatic(name, author).then((r) => {
			console.log(r);
			if (r.status == 200) {
				navigate("/matic/" + r.data);
			}
		});
	}

	return (
		<div className="content-center h-screen">
			<div className="border-2 rounded-md w-[60%] p-5  border-slate-700 m-auto text-slate-700">
				<div className="text-3xl text-center">New Matic</div>
				<div className="grid grid-cols-1 mt-4">
					<Input
						placeholder="Name"
						onChange={(e) => {
							setname(e.target.value);
						}}
					></Input>
					<Input
						placeholder="Author"
						onChange={(e) => {
							setAuthor(e.target.value);
						}}
					></Input>
					<div className="flex">
						<div className="w-[100px] ms-auto me-3">
							<Button2
								onClick={(e) => {
									save(e);
								}}
							>
								{" "}
								<Save /> Save{" "}
							</Button2>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
}
