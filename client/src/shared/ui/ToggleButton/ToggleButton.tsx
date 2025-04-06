import { memo, useCallback, useState } from "react";
import type { FC, ReactNode } from "react";
import cn from "clsx";

export interface ToggleButtonProps {
	children: ReactNode;
	defaultToggle?: boolean;
	onClick?: (active: boolean) => boolean | undefined;
}

const ToggleButton: FC<ToggleButtonProps> = memo(
	({ children, defaultToggle = false, onClick = () => {} }) => {
		const [active, setActive] = useState<boolean>(defaultToggle);

		const handle = useCallback(() => {
			const click = onClick(active);
			if (click !== undefined) {
				setActive(click);
			}
		}, [onClick, active]);

		const classes = cn(
			active
				? "text-bg-tertiary bg-accent-primary hover:bg-accent-secondary"
				: "text-accent-primary bg-bg-tertiary hover:bg-bg-quaternary",
			"rounded-lg h-8 w-full flex items-center justify-center hover:cursor-pointer transition-colors duration-150 box-border",
		);

		return (
			<button type="button" className={classes} onClick={handle}>
				{children}
			</button>
		);
	},
);

export { ToggleButton };
