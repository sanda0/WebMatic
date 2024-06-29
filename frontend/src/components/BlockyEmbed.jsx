import { useEffect, useState } from "react";
import { BlocklyWorkspace } from "react-blockly";
import "./../lib/customBlocks";

export default function BlockyEmbed(props) {
	const toolboxCategories = {
		kind: "categoryToolbox",
		contents: [
			{
				kind: "category",
				name: "Control",
				colour: "#5CA65C",
				contents: [
					{
						kind: "block",
						type: "control_start",
					},
					{
						kind: "block",
						type: "control_navigate",
					},
				],
			},
			{
				kind: "category",
				name: "Element",
				colour: "#5B67A5",
				contents: [
					{
						kind: "block",
						type: "element_by_css_selector",
					},
					{
						kind: "block",
						type: "element_by_xpath_selector",
					},
				],
			},
			{
				kind: "category",
				name: "Actions",
				colour: "#925AA5",
				contents: [
					{
						kind: "block",
						type: "action_click",
					},
					{
						kind: "block",
						type: "action_write",
					},
					{
						kind: "block",
						type: "action_wait",
					},
				],
			},
		],
	};




	return (
		<>
			<BlocklyWorkspace
				className="w-full h-[89vh]"
				toolboxConfiguration={toolboxCategories}
				onJsonChange={(j) => {
					let a = JSON.stringify(j);
					props.onJsonChange(a);
				}}
				initialXml={props.xml_data}
				onXmlChange={(x) => {
					props.onChange(x);
				}}
			/>
		</>
	);
}
