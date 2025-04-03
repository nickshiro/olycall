import type { Meta, StoryObj } from "@storybook/react";

import { Member } from "./Member";

const meta: Meta<typeof Member> = {
	component: Member,
	title: "Entities/Member",
};

export default meta;
type Story = StoryObj<typeof Member>;

export const Small: Story = {
	args: {
		name: "Evgeny Ponasenkov",
		avatar: "https://m.dedkov.space/meme/ponasenkov",
		isMuted: true,
		isDeaf: true,
		isScreencast: true,
		isVideo: true,
		isActive: true,
	},
	decorators: [
		(Story) => (
			<div style={{ width: 342 }}>
				<Story />
			</div>
		),
	],
};
