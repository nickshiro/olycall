import type { Meta, StoryObj } from "@storybook/react";

import { VoiceAvatar } from "./VoiceAvatar";

const meta: Meta<typeof VoiceAvatar> = {
	component: VoiceAvatar,
	title: "Entities/VoiceAvatar",
};

export default meta;

type Story = StoryObj<typeof VoiceAvatar>;

export const Default: Story = {
	args: {
		src: "//m.dedkov.space/meme/ponasenkov",
		isActive: true,
	},
};
