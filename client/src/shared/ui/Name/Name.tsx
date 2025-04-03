import type { FC } from "react";
import { memo } from "react";

export interface NameProps {
	children: string;
}

const NameComponent: FC<NameProps> = ({ children }) => {
	return (
		<h2 className="text-text-primary text-base font-medium selection:bg-primary selection:text-black">
			{children}
		</h2>
	);
};

export const Name = memo(NameComponent);
