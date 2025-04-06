import type { Meta, StoryObj } from "@storybook/react";

import { Avatar } from "./Avatar";

const meta: Meta<typeof Avatar> = {
	component: Avatar,
	title: "Shared/Avatar",
};

export default meta;

type Story = StoryObj<typeof Avatar>;

export const Default: Story = {
	args: {
		src: "//m.dedkov.space/meme/ponasenkov",
		size: "large",
	},
};
