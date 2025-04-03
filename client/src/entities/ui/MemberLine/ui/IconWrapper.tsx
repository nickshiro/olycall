import type { ReactNode, FC } from "react";
import { memo } from "react";

export interface IconWrapperProps {
	children: ReactNode;
}

const IconWrapperComponent: FC<IconWrapperProps> = ({ children }) => {
	return <li className="h-5 w-5 text-primary">{children}</li>;
};

export const IconWrapper = memo(IconWrapperComponent);
