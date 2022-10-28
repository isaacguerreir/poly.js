import { GenbankJs } from 'genbank'

declare namespace genbank {
  function read(path: string): Promise<GenbankJs>
}
