import { memo } from "react";
import type { FC, ReactNode } from "react";
import cn from "clsx";

export interface ButtonProps {
	children: ReactNode;
	onClick?: () => void;
	className?: string;
}

const Button: FC<ButtonProps> = memo(
	({ children, onClick = () => {}, className }) => {
		const classes = cn(
			className,
			!className && "hover:bg-bg-tertiary",
			"h-8 text-accent-primary rounded-lg hover:cursor-pointer flex items-center justify-center transition-colors duration-150 box-border w-full",
		);

		return (
			<button type="button" onClick={onClick} className={classes}>
				{children}
			</button>
		);
	},
);

export { Button };
