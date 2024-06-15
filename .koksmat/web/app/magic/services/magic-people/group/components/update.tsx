/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
"use client";
// piratos
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import { GroupForm } from "./form";

import { GroupItem } from "../applogic/model";
export default function UpdateGroup(props: { id: number }) {
  const { id } = props;

  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<GroupItem>(
    "magic-people.group",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const group = readResult.data;
  return <div>{group && <GroupForm group={group} editmode="update" />}</div>;
}
