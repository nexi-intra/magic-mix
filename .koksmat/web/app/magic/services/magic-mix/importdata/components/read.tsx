    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
import { useService } from "@/koksmat/useservice";
import { useState } from "react";
import {ImportDataItem} from "../applogic/model";


/* yankiebar */

export default function ReadImportData(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ImportDataItem>(
      "magic-mix.importdata",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const importdata = readResult.data;
    return (
      <div>
          
    {importdata && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{importdata.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{importdata.description}</div>
    </div>                <div>
                    <div className="font-bold" >data</div>
                    <div>{JSON.stringify(importdata.data,null,2)}</div>
                </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{importdata.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{importdata.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{importdata.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{importdata.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{importdata.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
