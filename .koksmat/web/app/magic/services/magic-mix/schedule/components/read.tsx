    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import {ScheduleItem} from "../applogic/model";


/* yankiebar */

export default function ReadSchedule(props: { id: number }) {
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
      <div>
          
    {schedule && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{schedule.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{schedule.description}</div>
    </div>    <div>
        <div className="font-bold" >cron</div>
        <div>{schedule.cron}</div>
    </div>    <div>
        <div className="font-bold" >job</div>
        <div>{schedule.job_id}</div>
    </div>                <div>
                    <div className="font-bold" >data</div>
                    <div>{JSON.stringify(schedule.data,null,2)}</div>
                </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{schedule.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{schedule.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{schedule.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{schedule.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{schedule.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
