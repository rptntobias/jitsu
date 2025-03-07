import { modeParameter, tableName } from './common';
import { intType, stringType, passwordType, booleanType } from '../../sources/types';

const icon = (
  <svg
    viewBox="0 0 25.6 25.6"
    height="100%"
    width="100%"
    xmlns="http://www.w3.org/2000/svg"
  >
    <g fill="none" stroke="#fff">
      <path
        d="m18.983 18.636c.163-1.357.114-1.555 1.124-1.336l.257.023c.777.035 1.793-.125 2.4-.402 1.285-.596 2.047-1.592.78-1.33-2.89.596-3.1-.383-3.1-.383 3.053-4.53 4.33-10.28 3.227-11.687-3.004-3.84-8.205-2.024-8.292-1.976l-.028.005c-.57-.12-1.2-.19-1.93-.2-1.308-.02-2.3.343-3.054.914 0 0-9.277-3.822-8.846 4.807.092 1.836 2.63 13.9 5.66 10.25 1.109-1.334 2.179-2.461 2.179-2.461.53.353 1.167.533 1.834.468l.052-.044a2.01 2.01 0 0 0 .021.518c-.78.872-.55 1.025-2.11 1.346-1.578.325-.65.904-.046 1.056.734.184 2.432.444 3.58-1.162l-.046.183c.306.245.285 1.76.33 2.842s.116 2.093.337 2.688.48 2.13 2.53 1.7c1.713-.367 3.023-.896 3.143-5.81"
        fill="#000"
        stroke="#000"
        strokeWidth="2.149"
      />
      <path
        d="m23.535 15.6c-2.89.596-3.1-.383-3.1-.383 3.053-4.53 4.33-10.28 3.228-11.687-3.004-3.84-8.205-2.023-8.292-1.976l-.028.005a10.31 10.31 0 0 0 -1.929-.201c-1.308-.02-2.3.343-3.054.914 0 0-9.278-3.822-8.846 4.807.092 1.836 2.63 13.9 5.66 10.25 1.116-1.342 2.186-2.469 2.186-2.469.53.353 1.167.533 1.834.468l.052-.044a2.02 2.02 0 0 0 .021.518c-.78.872-.55 1.025-2.11 1.346-1.578.325-.65.904-.046 1.056.734.184 2.432.444 3.58-1.162l-.046.183c.306.245.52 1.593.484 2.815s-.06 2.06.18 2.716.48 2.13 2.53 1.7c1.713-.367 2.6-1.32 2.725-2.906.088-1.128.286-.962.3-1.97l.16-.478c.183-1.53.03-2.023 1.085-1.793l.257.023c.777.035 1.794-.125 2.39-.402 1.285-.596 2.047-1.592.78-1.33z"
        fill="#336791"
        stroke="none"
      />
      <g strokeWidth=".716">
        <g strokeLinecap="round">
          <path
            d="m12.814 16.467c-.08 2.846.02 5.712.298 6.4s.875 2.05 2.926 1.612c1.713-.367 2.337-1.078 2.607-2.647l.633-5.017m-8.922-14.615s-9.284-3.796-8.852 4.833c.092 1.836 2.63 13.9 5.66 10.25 1.106-1.333 2.106-2.376 2.106-2.376m6.1-13.4c-.32.1 5.164-2.005 8.282 1.978 1.1 1.407-.175 7.157-3.228 11.687"
            strokeLinejoin="round"
          />
          <path
            d="m20.425 15.17s.2.98 3.1.382c1.267-.262.504.734-.78 1.33-1.054.49-3.418.615-3.457-.06-.1-1.745 1.244-1.215 1.147-1.652-.088-.394-.69-.78-1.086-1.744-.347-.84-4.76-7.29 1.224-6.333.22-.045-1.56-5.7-7.16-5.782s-5.423 6.885-5.423 6.885"
            strokeLinejoin="bevel"
          />
        </g>
        <g strokeLinejoin="round">
          <path d="m11.247 15.768c-.78.872-.55 1.025-2.11 1.346-1.578.325-.65.904-.046 1.056.734.184 2.432.444 3.58-1.163.35-.49-.002-1.27-.482-1.468-.232-.096-.542-.216-.94.23z" />
          <path
            d="m11.196 15.753c-.08-.513.168-1.122.433-1.836.398-1.07 1.316-2.14.582-5.537-.547-2.53-4.22-.527-4.22-.184s.166 1.74-.06 3.365c-.297 2.122 1.35 3.916 3.246 3.733"
            strokeLinecap="round"
          />
        </g>
      </g>
      <g fill="#fff">
        <path
          d="m10.322 8.145c-.017.117.215.43.516.472s.558-.202.575-.32-.215-.246-.516-.288-.56.02-.575.136z"
          strokeWidth=".239"
        />
        <path
          d="m19.486 7.906c.016.117-.215.43-.516.472s-.56-.202-.575-.32.215-.246.516-.288.56.02.575.136z"
          strokeWidth=".119"
        />
      </g>
      <path
        d="m20.562 7.095c.05.92-.198 1.545-.23 2.524-.046 1.422.678 3.05-.413 4.68"
        strokeLinecap="round"
        strokeLinejoin="round"
        strokeWidth=".716"
      />
    </g>
  </svg>
);

const postgresDestination = {
  description: <>Postgres SQL is a one of the most popular databases. While it's not suitable for large datasets (more than 100m events),
  it's probably an easiest way to start collecting data with Jitsu<br/><br />
  Jitsu works with Postgres both in stream and batch mode</>,
  syncFromSourcesStatus: 'supported',
  id: 'postgres',
  type: 'database',
  displayName: 'Postgres SQL',
  hidden: false,
  ui: {
    icon: icon,
    connectCmd: (cfg: object) => {
      return `PGPASSWORD="${cfg['_formData']['pgpassword']}" psql -U ${cfg['_formData']['pguser']} -d ${cfg['_formData']['pgdatabase']} -h ${cfg['_formData']['pghost']} -p ${cfg['_formData']['pgport']} -c "SELECT 1"`
    },
    title: (cfg: object) => {
      return cfg['_formData']['pghost'];
    }
  },
  parameters: [
    {
      id: '$type',
      constant: 'PostgresConfig'
    },
    modeParameter('stream'),
    tableName(),
    {
      id: '_formData.pghost',
      displayName: 'Host',
      required: true,
      type: stringType
    },
    {
      id: '_formData.pgport',
      displayName: 'Port',
      required: true,
      defaultValue: 5432,
      type: intType
    },
    {
      id: '_formData.pgdatabase',
      displayName: 'Database',
      required: true,
      type: stringType
    },
    {
      id: '_formData.pgschema',
      displayName: 'Schema',
      defaultValue: 'public',
      required: true,
      type: stringType
    },
    {
      id: '_formData.pguser',
      displayName: 'Username',
      required: true,
      type: stringType
    },
    {
      id: '_formData.pgpassword',
      displayName: 'Password',
      required: true,
      type: passwordType
    },
    {
      id: '_formData.pgdisablessl',
      displayName: 'Disable SSL',
      required: true,
      type: booleanType,
      defaultValue: false,
      documentation: <>
        All connections to Postgres will be unsecured (non-SSL). We do not recommend to disable SSL. Disabled SSL can be used with Postgres that is installed on the local machine.
      </>
    }
  ]

} as const;

export default postgresDestination;
