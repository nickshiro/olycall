import { memo } from "react";
import type { FC } from "react";
import cn from "clsx";

type Sizes = "small" | "large";

export interface AvatarProps {
	src: string;
	size?: Sizes;
}

const Avatar: FC<AvatarProps> = memo(({ src, size = "small" }) => {
	const sizeClass: Record<Sizes, string> = {
		large: "w-20 h-20",
		small: "w-10 h-10",
	};

	const classes = cn(
		"bg-center bg-cover bg-no-repeat rounded-full",
		sizeClass[size],
	);

	return <div className={classes} style={{ backgroundImage: `url(${src})` }} />;
});

export { Avatar };
