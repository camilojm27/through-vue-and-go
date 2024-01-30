export interface RootObject {
  took: number;
  timed_out: boolean;
  _shards: Shards;
  hits: Hits;
}

interface Shards {
  total: number;
  successful: number;
  skipped: number;
  failed: number;
}

interface Hits {
  total: Total;
  max_score: number;
  hits: Hit[];
}

interface Total {
  value: number;
}

interface Hit {
  _index: string;
  _type: string;
  _id: string;
  _score: number;
  '@timestamp': string;
  _source: Source;
}

interface Source {
  '@timestamp': string;
  Body: string;
  CC: string;
  'Content-Transfer-Encoding': string;
  'Content-Type': string;
  Date: string;
  From: string;
  'Message-Id': string;
  'Mime-Version': string;
  Subject: string;
  To: string;
  'X-FileName': string;
  'X-Folder': string;
  'X-From': string;
  'X-Origin': string;
  'X-To': string;
  'X-bcc': string;
  'X-cc': string;
}
