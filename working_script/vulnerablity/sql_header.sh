while read -r line;
         do 
            /var/root/go/bin/sql_header $line "csrftoken=d4b0e54c81be1d4d852359aa2a0e24b7; _pinterest_sess=TWc9PSZ5dklEREJYelJ5alpjUkt4MWk2bGZCQ3JReTlPR29CNHJ2QnFYWkxicG5LcEJwSVZKT0NidStVcXpDd0dQU3NrK0tldzhORkV6ekdvZ2dZKzBBMnowdGhFNkI1VUdJWFVsbjZGZWNXVk5pMWRPV2Q3QVRCTk1ncFhNdUdSWjlpSExOWHR5QnpZeDFPVFFqUy9qcm5FYXgzU3g5RU1kdDZBem82UkdpQ3dibndyWVRCeFVpREZhd3JlRm8rZVg4WjJCVE1pdDVHanNIOUQzdy9EUXQ5TTQwUXFOenFGUE5hOTM4VVpBSzlnL2lZejBnbUhXNklyZUd2N2hOOXBybkJVNVhocUFSRVdwcjVRUWJjNmo1U0VMZUt3NjFwNmltZVI4K2x0YTlReUxTQ0ZJMzZDb3dDOGVDb054eENsL2dFRmErbjZzNUtlb2JFMzFoUTBnTmRNNnI3WjVBNmJ1MEkvNk1jVlc4VVJ3SUJnZ2hsR3BWbEtsNWdBZDRKWnpVdTBXVWUvUWg4UVRlQkNPd24rZUx1OFdzWktVRzN4dSs2NE04Qi8xMWY2bGtmbFVjLzExS2lTT2w5T1FDUGlmSVowSG9NYW11ZHBIRGFZcHpmYTRSQlU2eFBEMUNjZGl4aW5yakFtKys4dFduUmhTOHhvSjk2bThIYVNUbVEzWi9NTmJMQ3JVdDhFVUl5TmRTU2pmM2VlTCtvTFMycWluZmszWTBCUHo4MXRNc3FNTXNMbTdMSEJuSlBIL1Rha2dQMkUrQTh6UmhZOWJwWlNldkNPSzg2TFpYeXVmeFhaN0MvWCs0T1hMMkdyLzk3LzcweFBRRmwwUnluOWZYQTdBS0k0dVhMbXdoK0tIR3ZYVVhyaXdubkZhcXpQRTdpSGdWR3VKenE5czN1UHBFQ0doZnNrV2NHV0VHUGYvWkxSYmx3K2RhSDdYajlNbHZ3TFd3S21pNXdTK29pcTBaSmE4c0diTkpuMkMwcFUzZ2kxMlBYSlE2QVNBU25Sa0ZwRURzbXA2S25YRVdzMUJqb1NDUTNuaWNXZEQzTWUwQldRWUxabEZRengvcEM5MW1ndVlKRXZldml0SFAxenhROS84NER0JkRjb25CMkhTQ0hxTkNPU3JFeXp4RDNBUGZ1cz0=; _auth=1; _routing_id="f172143c-c5ed-4ab1-8e68-f1e763835560"; sessionFunnelEventLogged=1; bei=false; cm_sub=none"
         done < $1
