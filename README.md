# file watcher
### what is it?
A file watcher is a program that looks for any changes to files or directores automatically, in real-time. For example, you might want to use these when you want to know if anything was added, removed, or altered in a linode bucket.

### how does it?
REmote DIctionary Server (redis).
  1. open source, in-memory data structure store
  2. stored in RAM, not disk (super fast) aka in-memory
  3. supports hashes, store key:value (filename:hash of file)
  4. Supports pub/sub messaging!
  5. Persistence by using disk for durability

Use redis to store filenames and their hashes to see if hash has changed. If the hash of the file has changed, indicating there was a change, we tell the file watcher to publish a message to a topic using Google Cloud Console. This message can be later used by a subcriber who is listening on the topic and another program will take over from there to download the changed files.

<img width="477" alt="image" src="https://github.com/user-attachments/assets/1eb0cb22-b59e-4cee-b0bd-9a4a5bb332c2">

### creating module and running Go program
<img width="425" alt="image" src="https://github.com/user-attachments/assets/18e1e445-e1b0-4a22-b7dd-d6ddb63a8a70">
