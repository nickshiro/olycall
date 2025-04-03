import type { FC } from "react";

const ArrowUp: FC = () => {
	return (
		<svg viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
			<title>Arrow up</title>
			<path
				d="M10 16V4M10 4L4.5 9.25M10 4L15.5 9.25"
				stroke="currentColor"
				stroke-width="2"
				stroke-linecap="round"
				stroke-linejoin="round"
			/>
		</svg>
	);
};

export { ArrowUp };
