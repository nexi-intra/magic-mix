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
import {ScheduleForm} from "./form";

import {ScheduleItem} from "../applogic/model";
export default function UpdateSchedule(props: { id: number }) {
    const { id } = props;
 
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ScheduleItem>(
      "magic-mix.schedule",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const schedule = readResult.data;
    return (
      <div>{schedule && 
      <ScheduleForm schedule={schedule} editmode="update"/>}
     
      </div>
    );
}
