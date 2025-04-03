import { memo } from "react";
import type { FC } from "react";
import type { Icon as IIcon } from "./Icon.type";
import { icons } from "./model/Icons";

export interface IconProps {
	icon: IIcon;
}

const IconComponent: FC<IconProps> = ({ icon }) => {
	const Element = icons[icon];

	return <Element />;
};

export const Icon = memo(IconComponent);
