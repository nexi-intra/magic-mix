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
import { UserForm } from "./form";

import { UserItem } from "../applogic/model";
export default function UpdateUser(props: { id: number }) {
  const { id } = props;

  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<UserItem>(
    "magic-people.user",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const user = readResult.data;
  return <div>{user && <UserForm user={user} editmode="update" />}</div>;
}
