import { memo } from "react";
import type { FC } from "react";
import type { IIcons } from "./model/icons";
import { icons } from "./model/icons";

export interface IconProps {
	icon: IIcons;
}

const Icon: FC<IconProps> = memo(({ icon }) => {
	const Self = icons[icon];

	if (!Self) return;

	return (
		<div className="h-full w-full flex items-center justify-center">
			<Self />
		</div>
	);
});

export { Icon };
