import ApplicationServices from 'lib/services/ApplicationServices';
import Marshal from 'lib/commons/marshalling';
import { closeableMessage, handleError } from 'lib/components/components';
import { Tab } from 'ui/components/Tabs/TabsConfigurator';

const destinationEditorUtils = {
  testConnection: async (dst: DestinationData, hideMessage?: boolean) => {
    try {
      await ApplicationServices.get().backendApiClient.post(
        '/destinations/test',
        Marshal.toPureJson(dst)
      );

      dst._connectionTestOk = true;

      if (!hideMessage) {
        closeableMessage.info('Successfully connected!');
      }
    } catch (error) {
      dst._connectionTestOk = false;
      dst._connectionErrorMessage = error.message ?? 'Failed to connect';

      if (!hideMessage) {
        handleError(error, 'Connection failed');
      }
    }
  },
  getPromptMessage: (tabs: Tab[]) => () =>
    tabs.some((tab) => tab.touched)
      ? 'You have unsaved changes. Are you sure you want to leave the page?'
      : undefined
};

export { destinationEditorUtils };
