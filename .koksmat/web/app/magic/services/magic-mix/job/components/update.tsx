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
import {JobForm} from "./form";

import {JobItem} from "../applogic/model";
export default function UpdateJob(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<JobItem>(
      "magic-mix.job",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const job = readResult.data;
    return (
      <div>{job && 
      <JobForm job={job} editmode="update"/>}
     
      </div>
    );
}
