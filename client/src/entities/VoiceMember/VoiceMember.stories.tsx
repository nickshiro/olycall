import type { Meta, StoryObj } from "@storybook/react";

import { VoiceMember } from "./VoiceMember";

const meta: Meta<typeof VoiceMember> = {
	component: VoiceMember,
	title: "Entities/VoiceMember",
};

export default meta;

type Story = StoryObj<typeof VoiceMember>;

export const Default: Story = {
	args: {
		isDeaf: true,
		isMuted: true,
		isVideo: true,
		isScreenshare: true,
		src: "//m.dedkov.space/meme/ponasenkov",
	},
	decorators: [
		(Story) => (
			<div className="w-85.5">
				<Story />
			</div>
		),
	],
};
