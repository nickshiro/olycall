import type { FC } from "react";
import { memo } from "react";
import cn from "classnames";

type Size = "small" | "large";

export interface AvatarProps {
	size?: Size;
	src?: string;
	isActive?: boolean;
	alt?: string;
}

const AvatarComponent: FC<AvatarProps> = ({
	size = "small",
	src = "",
	isActive = false,
}) => {
	const sizes: Record<Size, string> = {
		small: "h-10 w-10",
		large: "h-20 w-20",
	};

	return (
		<div className={cn(sizes[size], "relative box-border select-none")}>
			<div
				style={{ backgroundImage: `url(${src})` }}
				className="w-full h-full rounded-full bg-center bg-no-repeat bg-cover"
			/>
			{isActive && (
				<div className="w-full h-full rounded-full absolute top-0 left-0 box-border border-bg-secondary border-solid border outline outline-primary" />
			)}
		</div>
	);
};

export const Avatar = memo(AvatarComponent);
