import classNames from "classnames";
import type { FC, ReactNode } from "react";
import { memo } from "react";

export interface ButtonProps {
	children: ReactNode;
	isActive?: boolean;
	onClick?: (isActive: boolean) => void;
}

const ButtonComponent: FC<ButtonProps> = ({
	children,
	isActive = false,
	onClick = () => {},
}) => {
	const classes = classNames(
		isActive
			? "bg-primary text-bg-tertiary"
			: "bg-bg-tertiary text-primary hover:bg-bg-quaternary",
		"h-8 rounded-lg w-full flex cursor-pointer transition-colors duration-200",
	);

	return (
		<button type="button" className={classes} onClick={() => onClick(isActive)}>
			<div className="w-5 h-5 m-auto">{children}</div>
		</button>
	);
};

export const Button = memo(ButtonComponent);
