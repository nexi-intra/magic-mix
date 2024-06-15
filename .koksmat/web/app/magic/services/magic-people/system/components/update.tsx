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
import { SystemForm } from "./form";

import { SystemItem } from "../applogic/model";
export default function UpdateSystem(props: { id: number }) {
  const { id } = props;

  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<SystemItem>(
    "magic-people.system",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const system = readResult.data;
  return (
    <div>{system && <SystemForm system={system} editmode="update" />}</div>
  );
}
