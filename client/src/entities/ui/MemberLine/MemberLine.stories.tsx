import type { Meta, StoryObj } from "@storybook/react";

import { MemberLine } from "./MemberLine";

const meta: Meta<typeof MemberLine> = {
	component: MemberLine,
	title: "Entities/MemberLine",
};

export default meta;
type Story = StoryObj<typeof MemberLine>;

export const Default: Story = {
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
