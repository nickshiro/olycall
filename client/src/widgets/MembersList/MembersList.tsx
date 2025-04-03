import { MemberLine } from "@/entities/ui";
import type { FC } from "react";
import { memo } from "react";
import { JoinLine } from "./ui/JoinLine";

const members = [
	{ id: "121212", avatar: "//m.dedkov.space/ponasenkov", name: "Evgeny" },
	{ id: "323", avatar: "//m.dedkov.space/ponasenkov", name: "Evgeny" },
	{ id: "534534", avatar: "//m.dedkov.space/ponasenkov", name: "Evgeny" },
];

export interface MembersListProps {}

const MembersListComponent: FC<MembersListProps> = () => {
	return (
		<div className="p-2 box-border rounded-2xl bg-bg-secondary gap-y-1 flex flex-col">
			{members.map((member) => (
				<MemberLine avatar={member.avatar} name={member.name} key={member.id} />
			))}
			<JoinLine />
		</div>
	);
};

export const MembersList = memo(MembersListComponent);
