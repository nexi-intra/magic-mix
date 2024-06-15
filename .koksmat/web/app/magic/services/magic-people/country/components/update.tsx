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
import { CountryForm } from "./form";

import { CountryItem } from "../applogic/model";
export default function UpdateCountry(props: { id: number }) {
  const { id } = props;

  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<CountryItem>(
    "magic-people.country",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const country = readResult.data;
  return (
    <div>{country && <CountryForm country={country} editmode="update" />}</div>
  );
}
