import type { FC } from "react";
import { memo } from "react";

export interface TitleProps {
	children: string;
}

const TitleComponent: FC<TitleProps> = ({ children }) => {
	return (
		<h2 className="text-text-primary text-base font-medium selection:bg-primary selection:text-black">
			{children}
		</h2>
	);
};

export const Title = memo(TitleComponent);
