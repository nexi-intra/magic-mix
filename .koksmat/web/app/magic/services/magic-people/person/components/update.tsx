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
import { PersonForm } from "./form";

import { PersonItem } from "../applogic/model";
export default function UpdatePerson(props: { id: number }) {
  const { id } = props;

  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<PersonItem>(
    "magic-people.person",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const person = readResult.data;
  return (
    <div>{person && <PersonForm person={person} editmode="update" />}</div>
  );
}
