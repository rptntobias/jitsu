import {stringType, descriptionType, booleanType} from '../../sources/types';

const icon = <svg fill="none" version="1.1" height={"100%"} width={"100%"} viewBox="0 0 90 90" xmlns="http://www.w3.org/2000/svg"><path d="m86.175 3.7432c2.1157 2.0344 3.4991 4.7197 3.8246 7.6492 0 1.2206-0.3255 2.0344-1.0579 3.5805-0.7323 1.5461-9.7649 17.17-12.45 21.483-1.5461 2.5226-2.3599 5.5335-2.3599 8.4629 0 3.0109 0.8138 5.9404 2.3599 8.463 2.6853 4.3128 11.718 20.018 12.45 21.564 0.7324 1.5462 1.0579 2.2785 1.0579 3.4991-0.3255 2.9295-1.6275 5.6149-3.7432 7.5679-2.0344 2.1157-4.7197 3.4991-7.5678 3.7432-1.2206 0-2.0344-0.3255-3.4991-1.0579-1.4648-0.7324-17.414-9.5208-21.727-12.206-0.3255-0.1628-0.651-0.4069-1.0578-0.5697l-21.32-12.613c0.4882 4.0687 2.2785 7.9747 5.2079 10.823 0.5697 0.5696 1.1393 1.0579 1.7903 1.5461-0.4883 0.2441-1.0579 0.4883-1.5461 0.8138-4.3129 2.6853-20.018 11.718-21.564 12.45-1.5461 0.7324-2.2785 1.0579-3.5805 1.0579-2.9295-0.3255-5.6148-1.6275-7.5678-3.7432-2.1157-2.0344-3.4991-4.7197-3.8246-7.6492 0.081374-1.2206 0.40687-2.4412 1.0579-3.4991 0.73237-1.5461 9.7649-17.251 12.45-21.564 1.5461-2.5226 2.3599-5.4521 2.3599-8.4629 0-3.0109-0.8138-5.9404-2.3599-8.463-2.6853-4.4755-11.799-20.181-12.45-21.727-0.651-1.0579-0.9765-2.2785-1.0579-3.4991 0.3255-2.9295 1.6275-5.6148 3.7432-7.6492 2.0344-2.1157 4.7197-3.4177 7.6492-3.7432 1.2206 0.081374 2.4412 0.40687 3.5805 1.0579 1.302 0.56962 12.776 7.2423 18.879 10.823l1.3834 0.8137c0.4882 0.3255 0.8951 0.5696 1.2206 0.7324l0.651 0.4068 21.727 12.857c-0.4882-4.8825-3.0108-9.3581-6.9168-12.369 0.4883-0.2441 1.0579-0.4883 1.5461-0.8138 4.3129-2.6853 20.018-11.799 21.564-12.45 1.0579-0.651 2.2785-0.9765 3.5805-1.0579 2.8481 0.3255 5.5335 1.6275 7.5678 3.7432zm-40.036 47.034 4.6384-4.6384c0.651-0.651 0.651-1.6274 0-2.2784l-4.6384-4.6384c-0.651-0.651-1.6274-0.651-2.2784 0l-4.6384 4.6384c-0.651 0.651-0.651 1.6274 0 2.2784l4.6384 4.6384c0.5696 0.5696 1.6274 0.5696 2.2784 0z" fill="#ff694a"/></svg>

const dbtcloudDestination = {
    description: <>
        Special destination. The purpose of this destination is to trigger <b>dbt Cloud</b> Job on successful run of linked Connectors.
        See <a href={"https://docs.getdbt.com/dbt-cloud/api#operation/triggerRun"}>dbt Cloud API Docs</a>
        All other types of events are ignored.
    </>,
    syncFromSourcesStatus: 'not_supported',
    id: 'dbtcloud',
    type: 'other',
    displayName: 'dbt Cloud',
    hidden: true,
    ui: {
        icon,
        title: (cfg) => `Account ID: ${cfg._formData.dbtAccountId} Job ID: ${cfg._formData.dbtJobId}`,
        connectCmd: (_) => null
    },
    parameters: [
        {
            id: '_formData.description',
            displayName: 'Description',
            required: false,
            type: descriptionType,
            defaultValue: (
                <span>
                    Setup triggering of <b>dbt Cloud</b> Job on successful run of Sources and Destinations in batch mode.
                    <br />
                    See {' '}
                    <a target="_blank" href="https://docs.getdbt.com/dbt-cloud/api#operation/triggerRun">
                        dbt Cloud API Docs
                    </a>.
                </span>
            )
        },
        {
            id: '_formData.dbtEnabled',
            displayName: 'Enabled',
            defaultValue: false,
            required: true,
            type: booleanType
        },
        {
            id: '_formData.dbtAccountId',
            displayName: 'Account ID',
            required: true,
            type: stringType,
            documentation: <>
                Numeric ID of the dbt Cloud Account that the Job belongs to
            </>
        },
        {
            id: '_formData.dbtJobId',
            displayName: 'Job ID',
            required: true,
            type: stringType,
            documentation: <>
                Numeric ID of the Job to run
            </>
        },
        {
            id: '_formData.dbtCause',
            displayName: 'Cause',
            defaultValue: '`${_.event_type} ID: ${_.source}`',
            required: true,
            type: stringType,
            documentation: <>
                A text description of the reason for running this job.
                The value is treated as <a href={"https://jitsu.com/docs/configuration/javascript-functions"}>JavaScript functions</a>
            </>
        },
        {
            id: '_formData.dbtToken',
            displayName: 'Token',
            required: true,
            type: stringType,
            documentation: <>
                API Key
            </>
        },
    ]
}  as const;

export default dbtcloudDestination;
