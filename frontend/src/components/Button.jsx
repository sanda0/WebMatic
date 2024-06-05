export default function Button(props) {
	return (
		<button className="text-white flex p-2 m-1 gap-2 border w-full text-center rounded-md hover:bg-white hover:text-slate-700 ">
			{props.children}
		</button>
	);
}
