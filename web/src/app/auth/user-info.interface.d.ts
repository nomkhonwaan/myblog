/**
 * An authenticated user information
 */
interface UserInfo {
  /**
   * Name of the user
   */
  name: string;

  /**
   * Link to the user picture
   */
  picture: string;

  /**
   * Email of the user, will be presented only verified user
   */
  email: string;

  /**
   * A unique string generated by Auth0
   */
  sub: string;
}
