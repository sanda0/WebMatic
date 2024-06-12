import { useState } from "react";
import { BlocklyWorkspace } from "react-blockly";

export default function BlockyEmbed() {
	const toolboxCategories = {
		kind: "categoryToolbox",
		contents: [
			{
				kind: "category",
				name: "Logic",
				colour: "#5C81A6",
				contents: [
					{
						kind: "block",
						type: "controls_if",
					},
					{
						kind: "block",
						type: "logic_compare",
					},
				],
			},
			{
				kind: "category",
				name: "Math",
				colour: "#5CA65C",
				contents: [
					{
						kind: "block",
						type: "math_round",
					},
					{
						kind: "block",
						type: "math_number",
					},
				],
			},
			{
				kind: "category",
				name: "Custom",
				colour: "#5CA699",
				contents: [
					{
						kind: "block",
						type: "new_boundary_function",
					},
					{
						kind: "block",
						type: "return",
					},
				],
			},
		],
	};
	const initialXml =
		'<xml xmlns="http://www.w3.org/1999/xhtml"><block type="text" x="70" y="30"><field name="TEXT"></field></block></xml>';

	const [xml, setXml] = useState(
		'<xml xmlns="http://www.w3.org/1999/xhtml"><block type="text" x="70" y="30"><field name="TEXT"></field></block></xml>'
	);

	return (
		<BlocklyWorkspace
			className="w-full h-[89vh]" // you can use whatever classes are appropriate for your app's CSS
			toolboxConfiguration={toolboxCategories} // this must be a JSON toolbox definition
			initialXml={xml}
			onXmlChange={setXml}
		/>
	);
}
