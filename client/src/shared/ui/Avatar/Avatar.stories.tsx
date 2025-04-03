import type { Meta, StoryObj } from "@storybook/react";

import { Avatar } from "./Avatar";

const meta: Meta<typeof Avatar> = {
	component: Avatar,
	title: "Shared/Avatar",
};

export default meta;
type Story = StoryObj<typeof Avatar>;

export const Small: Story = {
	args: {
		size: "small",
		alt: "Evgeny",
		isActive: true,
		src: "https://m.dedkov.space/meme/ponasenkov",
	},
};

export const Large: Story = {
	args: {
		size: "large",
		alt: "Evgeny",
		isActive: false,
		src: "https://m.dedkov.space/meme/ponasenkov",
	},
};
