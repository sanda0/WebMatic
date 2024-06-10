export default function Input(props) {
	return (
		<input
			className="p-2 m-2 border rounded-md border-slate-700 focus:border-3 focus:border-slate-700 "
			onChange={(e) => {
				props.onChange(e);
			}}
			placeholder={props.placeholder}
		/>
	);
}
