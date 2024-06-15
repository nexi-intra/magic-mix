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
import {MapperItem} from "../applogic/model";


/* yankiebar */

export default function ReadMapper(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<MapperItem>(
      "magic-mix.mapper",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const mapper = readResult.data;
    return (
      <div>
          
    {mapper && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{mapper.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{mapper.description}</div>
    </div>    <div>
        <div className="font-bold" >source</div>
        <div>{mapper.source_id}</div>
    </div>    <div>
        <div className="font-bold" >transformation</div>
        <div>{mapper.transformation_id}</div>
    </div>    <div>
        <div className="font-bold" >target</div>
        <div>{mapper.target_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{mapper.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{mapper.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{mapper.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{mapper.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{mapper.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
