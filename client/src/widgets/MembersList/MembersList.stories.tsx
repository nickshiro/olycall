import type { Meta, StoryObj } from "@storybook/react";

import { MembersList } from "./MembersList";

const meta: Meta<typeof MembersList> = {
	component: MembersList,
	title: "Widgets/MembersList",
};

export default meta;
type Story = StoryObj<typeof MembersList>;

export const Default: Story = {
	args: {},
	decorators: [
		(Story) => (
			<div style={{ width: 358 }}>
				<Story />
			</div>
		),
	],
};
