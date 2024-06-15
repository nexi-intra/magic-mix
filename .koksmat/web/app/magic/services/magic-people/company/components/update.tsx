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
import { CompanyForm } from "./form";

import { CompanyItem } from "../applogic/model";
export default function UpdateCompany(props: { id: number }) {
  const { id } = props;

  const [transactionId, settransactionId] = useState(0);
  const readResult = useService<CompanyItem>(
    "magic-people.company",
    ["read", id?.toString()],
    "",
    6000,
    transactionId.toString()
  );
  const company = readResult.data;
  return (
    <div>{company && <CompanyForm company={company} editmode="update" />}</div>
  );
}
