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
import { RelationForm } from "./form";

import { RelationItem } from "../applogic/model";
export default function UpdateRelation(props: { id: number }) {
  const { id } = props;

  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<RelationItem>(
    "magic-people.relation",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const relation = readResult.data;
  return (
    <div>
      {relation && <RelationForm relation={relation} editmode="update" />}
    </div>
  );
}
